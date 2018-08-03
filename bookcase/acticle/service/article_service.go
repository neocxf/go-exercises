package service

import (
	"context"
	"github.com/neocxf/go-exercises/bookcase/acticle"
	"github.com/neocxf/go-exercises/bookcase/author"
	"github.com/neocxf/go-exercises/bookcase/models"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type articleService struct {
	articleRepos   article.ArticleRepository
	authorRepo     author.AuthorRepository
	contextTimeout time.Duration
}

type authorChannel struct {
	Author *models.Author
	Error  error
}

func NewArticleService(a article.ArticleRepository, ar author.AuthorRepository, time time.Duration) article.ArticleService {
	return &articleService{
		articleRepos:   a,
		authorRepo:     ar,
		contextTimeout: time,
	}
}

func (service *articleService) getAuthorDetail(ctx context.Context, item *models.Article, authorChan chan authorChannel) {
	res, err := service.authorRepo.GetById(ctx, item.Author.ID)
	holder := authorChannel{
		Author: res,
		Error:  err,
	}

	if ctx.Err() != nil {
		return // to avoid send on closed channel
	}

	authorChan <- holder
}

func (service *articleService) getAuthorDetails(ctx context.Context, data []*models.Article) ([]*models.Article, error) {
	chAuthor := make(chan authorChannel)

	defer close(chAuthor)

	existingAuthorMap := make(map[int64]bool)

	for _, item := range data {
		if _, ok := existingAuthorMap[item.Author.ID]; !ok {
			existingAuthorMap[item.Author.ID] = true
			go service.getAuthorDetail(ctx, item, chAuthor)
		}
	}

	mapAuthor := make(map[int64]*models.Author)
	totalGoroutineCalled := len(existingAuthorMap)

	for i := 0; i < totalGoroutineCalled; i++ {
		select {
		case a := <-chAuthor:
			if a.Error == nil {
				if a.Author != nil {
					mapAuthor[a.Author.ID] = a.Author
				}
			} else {
				return nil, a.Error
			}
		case <-ctx.Done():
			logrus.Warn("Timeout when calling user detail")
			return nil, ctx.Err()

		}
	}

	for index, item := range data {
		if a, ok := mapAuthor[item.Author.ID]; ok {
			data[index].Author = *a
		}
	}

	return data, nil
}

func (service *articleService) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Article, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	listArticle, err := service.articleRepos.Fetch(ctx, cursor, num)

	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	listArticle, err = service.getAuthorDetails(ctx, listArticle)

	if err != nil {
		return nil, "", err
	}

	if size := len(listArticle); size == int(num) {
		lastId := listArticle[num-1].ID
		nextCursor = strconv.Itoa(int(lastId))
	}

	return listArticle, nextCursor, nil
}

func (service *articleService) GetById(ctx context.Context, id int64) (*models.Article, error) {

	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	res, err := service.articleRepos.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	resAuthor, err := service.authorRepo.GetById(ctx, res.Author.ID)

	if err != nil {
		return nil, err
	}

	res.Author = *resAuthor

	return res, nil
}

func (service *articleService) Update(ctx context.Context, ar *models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()

	return service.articleRepos.Update(ctx, ar)
}

func (service *articleService) GetByTitle(ctx context.Context, title string) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	res, err := service.articleRepos.GetByTitle(ctx, title)

	if err != nil {
		return nil, err
	}

	resAuthor, err := service.authorRepo.GetById(ctx, res.Author.ID)

	if err != nil {
		return nil, err
	}

	res.Author = *resAuthor

	return res, nil
}

func (service *articleService) Store(ctx context.Context, ar *models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	existedArticle, _ := service.GetByTitle(ctx, ar.Title)

	if existedArticle != nil {
		return nil, models.CONFLICT_ERROR
	}

	id, err := service.articleRepos.Store(ctx, ar)

	if err != nil {
		return nil, err
	}

	ar.ID = id

	return ar, nil
}

func (service *articleService) Delete(ctx context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	existedArticle, _ := service.GetById(ctx, id)

	if existedArticle != nil {
		return false, models.NOT_FOUND_ERROR
	}

	return service.articleRepos.Delete(ctx, id)
}

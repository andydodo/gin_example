package linkservice

import (
	"regexp"

	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/pkg/errors"
)

const urlRe = `^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`

type LinkService struct {
	r types.LinkRepository
}

func New(linkRepository types.LinkRepository) *LinkService {
	return &LinkService{
		r: linkRepository,
	}
}

func (lS *LinkService) CreateLink(link *types.Link, url string) (*types.Link, error) {
	_, err := lS.r.FindByUserName(link.UserName)
	if err == nil {
		return &types.Link{}, errors.New("username already exists")
	}

	//add by andy
	ok, err := regexp.MatchString(urlRe, url)
	if !ok || err != nil {
		return &types.Link{}, errors.Wrap(err, "error url is illgal")
	}
	link.Url = url

	err = lS.r.Store(link)
	if err != nil {
		return &types.Link{}, errors.Wrap(err, "error storing link")
	}
	return link, nil
}

func (lS *LinkService) GetLink(id string) (*types.Link, error) {
	return lS.r.Find(id)
}

func (lS *LinkService) UpdateLink(link *types.Link) error {
	return lS.r.Update(link)
}

func (lS *LinkService) DeleteLink(id string) error {
	return lS.r.Delete(id)
}

func (lS *LinkService) GetAllLink(links *[]types.Link) error {
	return lS.r.FindAll(links)
}

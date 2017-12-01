package gorpo

import "image"

type ImageService struct {
	downloader Downloader
	repository EffectRepository
}

func NewImageService(downloader Downloader, repo EffectRepository) *ImageService {
	return &ImageService{
		downloader: downloader,
		repository: repo,
	}
}

func (i *ImageService) Process(imgSrc string, filters []Filter) (image.Image, error) {
	img, _, err := i.downloader.DownloadImage(imgSrc)

	if err != nil {
		return nil, err
	}

	for _, filter := range filters {
		effect, err := i.repository.GetEffect(filter.Id)

		if err != nil {
			return nil, err
		}

		img, err = effect.Transform(img, filter.Parameters)

		if err != nil {
			return nil, err
		}
	}

	return img, nil
}

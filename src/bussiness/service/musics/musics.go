package musics

import (
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain/musics"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
)

type mscSvs struct {
	mscDOm musics.MusicsItf
}

func (u *mscSvs) GetMusic() []model.Song {
	panic("implement me")
}

func (u *mscSvs) SaveFavorite() []model.FavoritesSongs {
	panic("implement me")
}

type MusicSvcItf interface {
	GetMusic()[]model.Song
	SaveFavorite()[]model.FavoritesSongs
}

func InitUserServices(msc musics.MusicsItf) MusicSvcItf {
	return &mscSvs{
		mscDOm: msc,
	}
}

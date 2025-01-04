package repository

import "vgstack-cli/templates/backend/db/model"

type LanguageRepository interface {
	AddLanguage(lang *model.Language) error
	DeleteLanguage(langId int) error
}

// var _ LanguageRepository = (*PostgreService)(nil)

func (p *PostgreService) AddLanguage(lang *model.Language) error {
	res, err := p.DB.Exec(`INSERT INTO "language"(name, release_year) VALUES($1, $2)`, lang.Name, lang.ReleaseYear)
	if err != nil {
		return err
	}
	langId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	lang.Id = int(langId)
	return nil
}

func (p *PostgreService) DeleteLanguage(langId int) error {
	_, err := p.DB.Exec(`DELETE FROM "language" WHERE id=$1`, langId)
	if err != nil {
		return err
	}
	return nil
}

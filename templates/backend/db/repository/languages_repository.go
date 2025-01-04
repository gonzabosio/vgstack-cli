package repository

import "vgstack-cli/templates/backend/db/model"

type LanguageRepository interface {
	AddLanguage(lang *model.Language) error
	DeleteLanguage(langId int) error
}

// var _ LanguageRepository = (*PostgreService)(nil)

func (p *PostgreService) AddLanguage(lang *model.Language) error {
	var langId int
	err := p.DB.QueryRow(`INSERT INTO "language"(name, release_year) VALUES($1, $2) RETURNING id`, lang.Name, lang.ReleaseYear).Scan(&langId)
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

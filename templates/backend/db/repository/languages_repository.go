package repository

import "vgstack-cli/templates/backend/db/model"

type LanguageRepository interface {
	AddLanguage(lang *model.Language) error
	ReadLanguages() ([]model.Language, error)
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

func (p *PostgreService) ReadLanguages() ([]model.Language, error) {
	var langs []model.Language
	rows, err := p.DB.Query(`SELECT * FROM language`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var lang model.Language
		if err := rows.Scan(&lang.Id, &lang.Name, &lang.ReleaseYear); err != nil {
			return nil, err
		}
		langs = append(langs, lang)
	}
	return langs, nil
}

func (p *PostgreService) DeleteLanguage(langId int) error {
	_, err := p.DB.Exec(`DELETE FROM "language" WHERE id=$1`, langId)
	if err != nil {
		return err
	}
	return nil
}

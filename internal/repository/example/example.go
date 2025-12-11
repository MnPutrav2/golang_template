package repository

func (q *exampleRepository) ExampleRepo(id string) error {

	if _, err := q.db.Exec("INSERT INTO test VALUES($1)", id); err != nil {
		return err
	}

	return nil
}
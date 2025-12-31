package main

type Directory struct {
	ID  int
	Dir string
}

func (a *App) GetLastSavedDir() ([]Directory, error) {
	rows, err := DB.Query("SELECT dir FROM last_dir_tb LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var directories []Directory
	for rows.Next() {
		var d Directory
		if err := rows.Scan(&d.Dir); err != nil {
			return nil, err
		}
		directories = append(directories, d)
	}

	return directories, nil
}

func (a *App) AddLastDir(dir string, id int) error {
	_, err := DB.Exec(
		"UPDATE last_dir_tb WHERE id = ? SET dir = ?",
		dir,
	)
	return err
}

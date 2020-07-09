package models

import (
	"context"
	"fmt"
)

type PostingPreview struct {
	Id string
	Title string
	Price int
	Photos []string
}

type PostingInput struct {
	Id string
	Title string
	Price int
	Condition string
	Description string
	Phone int
	City string
	Photos []string
	UserId string
}

func AllPostings() ([]*PostingPreview, error) {
	rows, err := db.Pool.Query(context.Background(),
		"SELECT id, title, price, photos FROM posting")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ps := make([]*PostingPreview, 0)
	for rows.Next() {
		p := new(PostingPreview)
		err := rows.Scan(&p.Id, &p.Title, &p.Price, &p.Photos)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

func CreatePosting(postingInput PostingInput) error {
	//err := db.Pool.Exec(context.Background(), "")
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO public.posting (title, price, condition, description, phone, city, photos, \"userId\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		postingInput.Title, postingInput.Price, postingInput.Condition, postingInput.Description,
		postingInput.Phone, postingInput.City, postingInput.Photos, postingInput.UserId)

	if err != nil {
		fmt.Errorf("%v", err)
		return err
	}

	return nil
}
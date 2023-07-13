// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: course.sql

package DB

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO Course (
    course_name,
    semester,
    description,
    u_id
) VALUES (
             $1, $2, $3, $4
         )
RETURNING c_id, course_name, semester, description, u_id
`

type CreateCourseParams struct {
	CourseName  string `json:"course_name"`
	Semester    string `json:"semester"`
	Description string `json:"description"`
	UID         int64  `json:"u_id"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse,
		arg.CourseName,
		arg.Semester,
		arg.Description,
		arg.UID,
	)
	var i Course
	err := row.Scan(
		&i.CID,
		&i.CourseName,
		&i.Semester,
		&i.Description,
		&i.UID,
	)
	return i, err
}

const getAllCourses = `-- name: GetAllCourses :many
SELECT c_id, course_name, semester, description, u_id from Course
`

func (q *Queries) GetAllCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getAllCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Course{}
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.CID,
			&i.CourseName,
			&i.Semester,
			&i.Description,
			&i.UID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourse = `-- name: GetCourse :one
SELECT c_id, course_name, semester, description, u_id from Course
WHERE c_id = $1
LIMIT 1
`

func (q *Queries) GetCourse(ctx context.Context, cID int64) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, cID)
	var i Course
	err := row.Scan(
		&i.CID,
		&i.CourseName,
		&i.Semester,
		&i.Description,
		&i.UID,
	)
	return i, err
}

const getCourseByName = `-- name: GetCourseByName :one
SELECT c_id, course_name, semester, description, u_id from Course
WHERE course_name = $1
LIMIT 1
`

func (q *Queries) GetCourseByName(ctx context.Context, courseName string) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourseByName, courseName)
	var i Course
	err := row.Scan(
		&i.CID,
		&i.CourseName,
		&i.Semester,
		&i.Description,
		&i.UID,
	)
	return i, err
}

const getCoursesBySemester = `-- name: GetCoursesBySemester :many
SELECT c_id, course_name, semester, description, u_id from Course
WHERE semester = $1
ORDER BY c_id
`

func (q *Queries) GetCoursesBySemester(ctx context.Context, semester string) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCoursesBySemester, semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Course{}
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.CID,
			&i.CourseName,
			&i.Semester,
			&i.Description,
			&i.UID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCoursesByUniversity = `-- name: GetCoursesByUniversity :many
SELECT c_id, course_name, semester, description, u_id from Course
WHERE u_id = $1
ORDER BY u_id
`

func (q *Queries) GetCoursesByUniversity(ctx context.Context, uID int64) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCoursesByUniversity, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Course{}
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.CID,
			&i.CourseName,
			&i.Semester,
			&i.Description,
			&i.UID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

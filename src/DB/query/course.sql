-- name: CreateCourse :one
INSERT INTO Course (
    course_name,
    semester,
    description,
    u_id
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetAllCourses :many
SELECT * from Course;

-- name: GetCourse :one
SELECT * from Course
WHERE c_id = $1
LIMIT 1;

-- name: GetCourseByName :one
SELECT * from Course
WHERE course_name = $1
LIMIT 1;

-- name: GetCoursesBySemester :many
SELECT * from Course
WHERE semester = $1
ORDER BY c_id;

-- name: GetCoursesByUniversity :many
SELECT * from Course
WHERE u_id = $1
ORDER BY u_id;


-- name: CreateStudentCourses :one
INSERT INTO Student_Courses (
    s_id,
    c_id
) VALUES (
          $1, $2
)
RETURNING *;

-- name: GetStudentCourse :one
SELECT * FROM Student_Courses
WHERE s_id = $1 and c_id = $2
LIMIT 1;

-- name: GetCoursesOfAStudent :many
SELECT * FROM Student_Courses
WHERE s_id = $1 ;

-- name: GetStudentsOfACourse :many
SELECT * FROM Student_Courses
WHERE c_id = $1 ;

-- name: DeleteCoursesOfAStudent :exec
DELETE FROM Student_Courses
WHERE s_id = $1;

-- name: ResetUsers :exec
DELETE FROM users;

-- name: ResetRoles :exec
DELETE FROM roles;

-- name: ResetUsersRoles :exec
DELETE FROM users_roles;

-- name: ResetDegrees :exec
DELETE FROM degree_programs;

-- name: ResetYears :exec
DELETE FROM years;

-- name: ResetUsersPrograms :exec
DELETE FROM users_programs;

-- name: ResetCourses :exec
DELETE FROM courses;

-- name: ResetDegreesCourses :exec
DELETE FROM degrees_courses;

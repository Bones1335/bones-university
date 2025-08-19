# Bones University

This app is an imitation of a university site like Blackboard, Moodle, and Canvas that has three roles: Admin, Professor, and Student. The dashboard for each role should change depending on who is logged in to the site. Once logged in, an admin can manage student enrollments, add document models for auto-generated admin documents, validate student grades submitted by professors, and more. Professors can assign and view submitted homework and input and submit grades. Students can submit homework, print or download administrative files to be completed and uploaded, and select their internship wishes.

# TODO:

  - [x] CREATE User/Enrollment endpoint
  - [x] READ User endpoint
  - [x] UPDATE User endpoint
  - [ ] DELETE User endpoint

  - [x] CREATE Roles endpoint
    - [x] Role types: unset, admin, professor, student
    
  - [ ] READ Roles endpoint
    - [x] read user's role
  - [ ] UPDATE Roles endpoint
    - [x] update user's role if admin is signed in
  - [ ] DELETE Roles endpoint

  - [x] Login endpoint

  - [ ] Upload document model endpoint

  - [x] CREATE Degree endpoint
    - this is an admin endpoint
  - [x] READ Degree endpoint
    - this is an endpoint for everyone
  - [ ] UPDATE Degree endpoint
    - this is an admin endpoint
  - [ ] DELETE Degree endpoint
    - this is an admin endpoint

  - [x] CREATE Year endpoint
  - [ ] READ Year endpoint
  - [ ] UPDATE Year endpoint
  - [ ] DELETE Year endpoint

  - [x] CREATE Student_Program endpoint
  - [ ] READ Student_Program endpoint
  - [ ] UPDATE Student_Program endpoint
  - [ ] DELETE Student_Program endpoint

  - [x] CREATE Courses endpoint
  - [ ] READ Courses endpoint
  - [ ] UPDATE Courses endpoint
  - [ ] DELETE Courses endpoint
  
  - Courses need to be connected to their given degree
    - [x] CREATE Degrees_Courses endpoint
    - [ ] READ Degrees_Courses endpoint
    - [ ] UPDATE Degrees_Courses endpoint
    - [ ] DELETE Degrees_Courses endpoint

  - [x] CREATE Assignments endpoint
  - [ ] READ Assignments endpoint
  - [ ] UPDATE Assignments endpoint
  - [ ] DELETE Assignments endpoint

  - [ ] CREATE Student_Assignments endpoint
  - [ ] READ Student_Assignments endpoint
  - [ ] UPDATE Student_Assignments endpoint
  - [ ] DELETE Student_Assignments endpoint

  - [ ] CREATE Grades endpoint
  - [ ] READ Grades endpoint
  - [ ] UPDATE Grades endpoint
  - [ ] DELETE Grades endpoint

  - [ ] Link Grades, Courses, Assignments, Students, and Professors
    - [ ] One Professor linked to many Courses
    - [ ] Table with many assignments per Course
    - [ ] Table with Course ID, Assignment ID, Student ID, and Grade associated to that assignment

  - [ ] CREATE Internships endpoint
    - an internship needs to designate the number of spots they have available
  - [ ] READ Internships endpoint
  - [ ] UPDATE Internships endpoint
  - [ ] DELETE Internships endpoint

  - [ ] Connect one internship to many students

  - [ ] Collect internship and housing data to generate a payment report for each student's completed internship

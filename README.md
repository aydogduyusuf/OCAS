# OCAS
A brief description of our project  
As students, each semester we face these problems:  
Which course will be the most suitable for us in terms of our workload and study time.  
We don’t exactly know what we will face at the courses we took.  
It is hard to find course related valuable resources except the uploaded resources by the course lecturer. If we can’t understand from these resources(slides, lectures etc.) we need an experienced student to tell us how to study and from which source we will study.  
To overcome these problems, we came up with the idea of OCAS. Briefly, we provide students the opportunity to find an experienced student and help them during the semester. Basically, we have two end-users: students and mentors. Mentors are the experienced students who take the course one of the previous semesters and have success in it. They apply for our system to be the mentor. If they are accepted, they earn money and start to assist students by uploading useful resources, making recitations, giving them tips, answering their questions and marking important deadlines about the course on the calendar etc.  
Students register our system and select a subscription type, pay money and take courses. To solve the problem of deciding which course students should take, we thought of this idea:  a suggestion system based on the data gathered from the registered students will be developed. A questionnaire will be created and sent to the students to be filled by the students. The aim of this step is to analyze the data of the students and try to estimate the success rate at the end of the semester.  

Online Course Assistant System API  
A demo version of the graduation project of group OCAS. This API offers these features:  
1. Mentor and Student Registration Page:  
Mentor and student registration pages are implemented and they are completely functional since authentication is essential for using OCAS.   
2. Mentor Course Management Page:  
After login, mentors can track the courses that they mentor and update course content. This page is completely functional for mentors to add or delete the course contents that can be accessed by students who are registered for that course.  
3. Student Profile Page:  
	Students have a public profile page where they can change their avatar and their public name, a courses page where they can see all the courses they are registered at that semester or previous ones, a subscription page where they can manage their subscription plan.  
4. Student Subscription Types Pricing Information Page:  
	In the student profile, there is a page implemented where students can subscribe to OCAS with different subscription types and with different durations. The subscription plan determines the maximum number of courses that can be registered by the student for that semester. And at the same page students can get information about different subscription types.  
5. Mentor Balance Management Page:  
	In the mentor profile there is a page implemented where mentors can see their balance, update their bank account details and request withdrawal from OCAS to their current bank account by filling a form.  
6. Mentor Evaluation Button:  
	A mentor rating system inside the course page.  
7. Calendar With Deadlines:  
	A calendar is implemented in the course details page where the mentor of that course can add various events to the course calendar and students who are registered to that course can use that calendar to track their homework deadlines, recitations and exams.  

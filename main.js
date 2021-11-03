//import User from "./class";
const User = require("./class.js");//commún
const hitesh = new User("Seba","ssp013@gmail.com");
hitesh.enrollCourse("Math");
hitesh.enrollCourse("Biology");
hitesh.enrollCourse("Spanish");
let courses =hitesh.getCourselist();
courses.forEach( (c) => console.log(c)); 
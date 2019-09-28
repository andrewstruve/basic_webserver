# basic_webserver
A basic webserver project template written in GO

The purpose for creating this project is to share with people how I like to build sites using GO. There are other great ways to make a website but I really enjoy being to craft a site to my own requirements. 

This project is meant to be a template for my own and if you should enjoy it,a starting point for your personal projects.

The architechture of the project is a MVC (Model-View-Controller architecture:

main handler sets up all the routes (main.go)

The controller takes in the data and either sends back a response or renders data (controller.go)

The view renders all web pages based on the input provided to each page render (view.go)

<* Coming Soon *> The model handles data being stored for the application. This could be for local data storage or for data base handling. The controller can forward data onto the model. (model.go)

There are 3 pages:
*Home
*Test Data
*About

Home and About are what they say they are.
Test Data is used to test out different types of Request types when data is submitted or retrieved 
from a page.


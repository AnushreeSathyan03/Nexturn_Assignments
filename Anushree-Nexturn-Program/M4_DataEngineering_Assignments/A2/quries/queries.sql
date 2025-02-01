-- Q1: List the names of employees hired after January 1, 2021
SELECT Name
FROM Employees
WHERE HireDate > '2021-01-01';

-- Q2: Calculate the average salary of employees in each department
SELECT D.DepartmentName, AVG(E.Salary) AS AverageSalary
FROM Employees E
JOIN Departments D ON E.DepartmentID = D.DepartmentID
GROUP BY D.DepartmentName;

-- Q3: Find the department name where the total salary is the highest
SELECT D.DepartmentName
FROM Employees E
JOIN Departments D ON E.DepartmentID = D.DepartmentID
GROUP BY D.DepartmentName
ORDER BY SUM(E.Salary) DESC
LIMIT 1;

-- Q4: List all departments that currently have no employees assigned
SELECT D.DepartmentName
FROM Departments D
LEFT JOIN Employees E ON D.DepartmentID = E.DepartmentID
WHERE E.EmployeeID IS NULL;

-- Q5: Fetch all employee details along with their department names
SELECT E.EmployeeID, E.Name, E.Salary, E.HireDate, D.DepartmentName
FROM Employees E
JOIN Departments D ON E.DepartmentID = D.DepartmentID;

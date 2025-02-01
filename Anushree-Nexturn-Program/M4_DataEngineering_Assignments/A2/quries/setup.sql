-- Create Employees Table
CREATE TABLE Employees (
    EmployeeID INTEGER PRIMARY KEY,
    Name TEXT NOT NULL,
    DepartmentID INTEGER NOT NULL,
    Salary REAL NOT NULL,
    HireDate TEXT NOT NULL,
    FOREIGN KEY (DepartmentID) REFERENCES Departments(DepartmentID)
);

-- Create Departments Table
CREATE TABLE Departments (
    DepartmentID INTEGER PRIMARY KEY,
    DepartmentName TEXT NOT NULL
);

-- Insert Data into Departments Table
INSERT INTO Departments (DepartmentID, DepartmentName) VALUES
(101, 'HR'),
(102, 'IT'),
(103, 'Finance');

-- Insert Data into Employees Table
INSERT INTO Employees (EmployeeID, Name, DepartmentID, Salary, HireDate) VALUES
(1, 'Alice', 101, 70000, '2021-01-15'),
(2, 'Bob', 102, 60000, '2020-03-10'),
(3, 'Charlie', 101, 80000, '2022-05-20'),
(4, 'Diana', 103, 75000, '2019-07-25');

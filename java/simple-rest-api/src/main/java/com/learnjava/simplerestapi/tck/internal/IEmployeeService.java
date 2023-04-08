package com.learnjava.simplerestapi.tck.internal;

import com.learnjava.simplerestapi.tck.domain.Employee;

import java.util.List;

public interface IEmployeeService {
    List<Employee> getEmployee();
    Employee getEmployeeById(int id);
    void addEmployee(Employee employee);
    void updateEmployee(int id, Employee employee);
    void deleteEmployee(int id);
}

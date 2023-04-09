package com.learnjava.simplerestapi.service;

import com.learnjava.simplerestapi.tck.domain.Employee;
import com.learnjava.simplerestapi.tck.internal.IEmployeeRepository;
import com.learnjava.simplerestapi.tck.internal.IEmployeeService;
import lombok.AllArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
@AllArgsConstructor
public class EmployeeService implements IEmployeeService {

    // Repository Injection
    @Autowired
    private final IEmployeeRepository employeeRepository;

    // Method
    public List<Employee> getEmployee() {
        return employeeRepository.findAll();
    }
    public Employee getEmployeeById(int id) {
        return null;
    }
    public void addEmployee(Employee employee) {}
    public void updateEmployee(int id, Employee employee) {}
    public void deleteEmployee(int id) {}
}

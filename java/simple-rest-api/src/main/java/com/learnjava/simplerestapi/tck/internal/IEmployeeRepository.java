package com.learnjava.simplerestapi.tck.internal;

import com.learnjava.simplerestapi.tck.domain.Employee;
// JPARepository is an interface in Spring Data
// that provides a set of methods to perform CRUD (Create, Read, Update, Delete) operations on a database
// using Java Persistence API (JPA).
import org.springframework.data.jpa.repository.JpaRepository;

// Interface "IBookRepository" extends JpaRepository<EntityClass, Identifier atau Primary key>
public interface IEmployeeRepository extends JpaRepository<Employee, Long> {
}

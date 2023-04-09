package com.learnjava.simplerestapi.tck.domain;

import lombok.Getter;
import lombok.Setter;

import java.time.LocalDateTime;

import javax.persistence.*;

/*
* @Entity mendeklarasikan class "Employee" sebagai sebuah JPA enitity.
* Ini berarti JPA akan menggunakan field yang ada di dalam class untuk membuat kolom di database.
* */
@Entity

/*
* @Table menunjuk nama tabel yang akan dimap ke class "Employee".
* Atau, memberitahu ke JPA bahwa class "Employee" bernama Employee.tableName atau "employees" pada database.
* */
@Table(name = Employee.tableName)

// @Getter dan @Setter berfungsi untuk otomatis membuatkan getter dan setter oleh plugin Lombok.
@Getter
@Setter

public class Employee {

    /*
    * Field "tableName" bisa diakses secara "public"
    * dan "final" yang mana fieldnya tidak bisa dioverride maupun diubah.
    * "static" juga menandakan bahwa tableName ini adalah field dari Class, bukan dari Objek.
    * */
    public static final String tableName = "employees";

    // @Id menunjuk bahwa field "id" akan dipetakan kekolom database sebagai primary key.
    @Id

    /*
    * The "@GeneratedValue(strategy = GenerationType.AUTO)" is a Java annotation used in conjunction with JPA (Java Persistence API)
    * to automatically generate unique primary key values for entities (objects that are stored in a database table).
    * The "strategy" attribute specifies the type of algorithm used to generate the primary key values.
    * In this case, the "AUTO" strategy indicates that the persistence provider should choose the appropriate strategy
    * based on the underlying database and its capabilities.
    * This annotation is often used in conjunction with the "@Id" annotation,
    * which indicates that a field is the primary key of the entity.
    * */
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long  id;

    // @Column menunjuk sebuah field dimap ke kolom database.
    @Column(name = "name")
    private String name;
    @Column(name = "release_date")
    private LocalDateTime joinDate;

}

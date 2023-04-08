package com.learnjava.simplerestapi.tck.domain;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table(name = Title.tableName)
@Getter
@Setter
public class Title {
    public static final String tableName = "titles";

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    @Column(name = "title_name")
    private String titleName;
    @Column(name = "salary")
    private Double salary;

}

package com.playground.springcore.bean;

import com.playground.springcore.common.domain.User;
import org.springframework.context.annotation.Bean;

public class RenameBeanConfiguration {

    @Bean
    public User user1() {
        System.out.println("Create User 1");
        return new User();
    }

    @Bean
    public User user2() {
        System.out.println("Create User 2");
        return new User();
    }

    @Bean("helmi")
    public User renameBean() {
        System.out.println("Create renameBean");
        return new User();
    }
}

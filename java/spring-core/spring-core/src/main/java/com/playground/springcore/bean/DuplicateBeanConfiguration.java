package com.playground.springcore.bean;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class DuplicateBeanConfiguration {

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
}

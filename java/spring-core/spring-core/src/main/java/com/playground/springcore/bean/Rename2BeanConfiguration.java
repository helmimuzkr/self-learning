package com.playground.springcore.bean;

import com.playground.springcore.common.domain.User;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class Rename2BeanConfiguration {

    @Bean
    public User user1() {
        System.out.println("Create User 1 in Rename2BeanConfiguration class");
        return new User();
    }

    @Bean
    public User user2() {
        System.out.println("Create User 2 in Rename2BeanConfiguration class");
        return new User();
    }

    @Bean("muzakir")
    public User renameBean() {
        System.out.println("Create renameBean in Rename2BeanConfiguration class");
        return new User();
    }
}

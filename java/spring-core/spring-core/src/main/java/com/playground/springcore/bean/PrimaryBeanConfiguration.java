package com.playground.springcore.bean;

import com.playground.springcore.common.domain.User;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;

@Configuration
public class PrimaryBeanConfiguration {

    // Default Bean
    @Primary
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

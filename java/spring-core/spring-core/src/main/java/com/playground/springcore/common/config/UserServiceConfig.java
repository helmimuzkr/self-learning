package com.playground.springcore.common.config;

import com.playground.springcore.common.service.UserService;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class UserServiceConfig {

    @Bean
    public UserService userService() {
        return new UserService();
    }
}

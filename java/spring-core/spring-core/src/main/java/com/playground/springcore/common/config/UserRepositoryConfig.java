package com.playground.springcore.common.config;

import com.playground.springcore.common.repository.UserRepository;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class UserRepositoryConfig {

    @Bean
    public UserRepository userRepository1() {
        return new UserRepository();
    }

    @Bean
    public UserRepository userRepository2() {
        return new UserRepository();
    }
}

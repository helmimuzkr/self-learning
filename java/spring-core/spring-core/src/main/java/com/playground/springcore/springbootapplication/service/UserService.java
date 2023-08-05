package com.playground.springcore.springbootapplication.service;


import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class UserService {

    public void saveUser(String name) {
        log.info("Start save user service : " + name);
    }
}

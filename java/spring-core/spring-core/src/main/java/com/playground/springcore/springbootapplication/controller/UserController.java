package com.playground.springcore.springbootapplication.controller;

import com.playground.springcore.springbootapplication.service.UserService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class UserController {

    @Autowired
    private UserService userService;

    public void saveUserController() {
        log.info("Start save user controller");
        String name = "helmi";
        userService.saveUser(name);
    }
}

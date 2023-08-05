package com.playground.springcore.springbootapplication;


import com.playground.springcore.common.domain.User;
import com.playground.springcore.springbootapplication.controller.UserController;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.Bean;

// the application class that use annotation @SpringBootApplication will scan all configurations or components
// in the application class package. it's because the default @ComponentScan in @SpringBootApplication.
// i.e: this UserApplication will scan all component/configuration in this package.
@SpringBootApplication
public class UserApplication {

    @Bean
    User user() {
        return new User();
    }

    public static void main(String[] args) {

        // we can bootstrap or launch the application using SpringApplication
        // SpringApplication also will creating application context.
        // https://docs.spring.io/spring-boot/docs/current/api/org/springframework/boot/SpringApplication.html
        ConfigurableApplicationContext applicationContext = SpringApplication.run(UserApplication.class, args);

        UserController userController = applicationContext.getBean(UserController.class);
        System.out.println(userController);
    }
}

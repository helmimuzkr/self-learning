package com.playground.springcore;

import com.playground.springcore.component.UserController;
import com.playground.springcore.common.repository.UserRepository;
import com.playground.springcore.common.service.UserService;
import com.playground.springcore.componentscan.ComponentScanConfig;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class CompoenntAndComponentScanTest {

    @Test
    public void testComponentScan() {
        ConfigurableApplicationContext applicationContext = new AnnotationConfigApplicationContext(ComponentScanConfig.class);

        // the beans using @Configuration
        UserRepository userRepository = applicationContext.getBean("userRepository1", UserRepository.class);
        UserService userService = applicationContext.getBean("userService", UserService.class);
        Assertions.assertNotNull(userRepository);
        Assertions.assertNotNull(userService);

        // the beans using @Component
        UserController userController = applicationContext.getBean("userController", UserController.class);
        Assertions.assertNotNull(userController);
    }
}

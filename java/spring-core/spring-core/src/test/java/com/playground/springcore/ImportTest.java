package com.playground.springcore;

import com.playground.springcore.importconfiguration.MainConfig;
import com.playground.springcore.common.repository.UserRepository;
import com.playground.springcore.common.service.UserService;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class ImportTest {

    @Test
    void testImport() {
        ConfigurableApplicationContext applicationContext = new AnnotationConfigApplicationContext(MainConfig.class);

        UserRepository userRepository1 = (UserRepository) applicationContext.getBean("userRepository1");
        UserRepository userRepository2 = applicationContext.getBean("userRepository2", UserRepository.class);
        Assertions.assertNotSame(userRepository1, userRepository2);

        UserService userService = applicationContext.getBean("userService", UserService.class);
        Assertions.assertNotNull(userService);
    }
}

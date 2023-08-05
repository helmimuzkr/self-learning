package com.playground.springcore;

import com.playground.springcore.aware.AuthService;
import com.playground.springcore.common.config.AwareConfig;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class AwareTest {

    // https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/beans/factory/Aware.html
    // https://youtu.be/VM3rwdMBORY?t=12624

    @Test
    public void testAware() {
        ConfigurableApplicationContext applicationContext = new AnnotationConfigApplicationContext(AwareConfig.class);



        AuthService authService = applicationContext.getBean(AuthService.class);

        // check if the bean name and application context has already injected to AuthService when Spring create the Bean.
        // because we using @ComponentScan, the bean name is class name with pascalCase
        Assertions.assertEquals("authService", authService.getBeanName());
        Assertions.assertSame(applicationContext, authService.getApplicationContext());
    }
}

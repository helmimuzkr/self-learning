package com.playground.springcore;

import com.playground.springcore.applicationcontext.UserConfiguration;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class ApplicationContextTest {

    @Test
    void testApplicationContext() {
        // Create Spring IOC/Container/Application Context
        // The parameter is configuration class
        ApplicationContext context = new AnnotationConfigApplicationContext(UserConfiguration.class);

        // Test if the context not null, if not null then Application Context successfully created.
        Assertions.assertNotNull(context);
    }
}

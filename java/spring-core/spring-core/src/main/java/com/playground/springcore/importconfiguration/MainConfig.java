package com.playground.springcore.importconfiguration;

import com.playground.springcore.common.config.UserRepositoryConfig;
import com.playground.springcore.common.config.UserServiceConfig;
import org.springframework.context.annotation.Import;

@Import({
        UserRepositoryConfig.class,
        UserServiceConfig.class
})
public class MainConfig {
}

package com.playground.springcore.component;

import org.springframework.stereotype.Component;

// Jika menggunakan anotasi @Component lalu di scan menggunakan @ComponentScan,
// maka class tersebut akan langsung dibuatkan bean nya tanpa method seperti menggunakan @Bean

// Nama bean otomatis dari class nya.
// i.e: userController
@Component
public class UserController {
}

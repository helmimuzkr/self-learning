package com.playground.springcore.bean;

import com.playground.springcore.common.domain.User;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class CreateBeanConfiguration {

    // Bean yang dibuat akan disimpan ke dalam Spring Container/Application Context

    // Jika method ditambahkan annotation @Bean maka bisa dikatakan bahwa method tersebut membuat bean.
    // Nama bean yang dihasilkan nanti defaultnya diambil dari nama methodnya.
    @Bean
    public User userBean() {
        System.out.println("Create User bean");
        return new User();
    }
}

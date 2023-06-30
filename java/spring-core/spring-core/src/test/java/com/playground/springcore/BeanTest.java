package com.playground.springcore;

import com.playground.springcore.bean.*;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.NoUniqueBeanDefinitionException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class BeanTest {

    @Test
    void createBeanTest() {
        // Buat application context yang menampung bean nya.
        ApplicationContext context = new AnnotationConfigApplicationContext(CreateBeanConfiguration.class);

        // Ambil bean dari context yang sudah dibuat dengan me-assign class dari bean yang diinginkan.
        User user = context.getBean(User.class);
        Assertions.assertNotNull(user);

        // Akan mendapatkan bean yang sama, dikarenakan konsep pembuatan Bean menggunakan Singleton Pattern.
        User user1 = context.getBean(User.class);
        User user2 = context.getBean(User.class);
        Assertions.assertSame(user1, user2);
    }

    @Test
    void duplicateBeanTest() {
        ApplicationContext context = new AnnotationConfigApplicationContext(DuplicateBeanConfiguration.class);

        // Ketika di dalam container ada 2 Bean dengan tipe yang sama, lalu kita memanggil salah satu bean tersebut,
        // Maka dapat dipastikan terkena Exception.
        // Spring bingung ingin memberikan bean yang mana, karena disini saat meminta Bean ke Spring
        // kita hanya memberitahukan tipe atau class dari bean yang kita inginkan.
        Assertions.assertThrows(NoUniqueBeanDefinitionException.class, () -> {
            User user1 = context.getBean(User.class);
        });

        // Supaya tidak ambigu, kita bisa memanggilnya berdasarkan nama bean nya.
        // Caranya dengan menambahkan string nama bean pada parameter pertama.
        User user1 = context.getBean("user1", User.class);
        Assertions.assertNotNull(user1);
    }

    @Test
    void primaryBeanTest() {
        ApplicationContext context = new AnnotationConfigApplicationContext(PrimaryBeanConfiguration.class);

        // Ketika method bean ditambahkan annotation @Primary, maka spring akan membuat method itu sebagai default.
        // Jadi ketika ada 2 Bean dengan tipe yang sama lalu kita memanggilnya hanya menggunakan tipe nya saja,
        // maka yang akan diberikan oleh spring adalah bean dengan annotation @Primary.
        // Contoh kasus disini yaitu bean dengan nama user1 adalah Primary.
        User userPrimary = context.getBean(User.class);
        User user1 = context.getBean("user1", User.class);
        User user2 = context.getBean("user2", User.class);

        // Maka userPrimary dan user1 adalah objek yang sama, karena mengambil bean yang sama.
        // userPrimary == user1
        Assertions.assertSame(userPrimary, user1);
        // userPrimary != user2
        Assertions.assertNotSame(userPrimary, user2);
    }

    @Test
    public void renamingBeanTest() {
        ApplicationContext context = new AnnotationConfigApplicationContext(RenameBeanConfiguration.class, Rename2BeanConfiguration.class);

        // Nama bean seharusnya unique, karena bisa saja terdapat nama bean yang sama pada configuration class yang berbeda.
        // Ini disebabkan karena method pembuat bean nya sama.
        // Defaultnya nama bean di ambil dari nama method pembuatnya.

        User user1classOne = context.getBean("user1", User.class);
        User user2classOne = context.getBean("user2", User.class);

        User user1classTwo = context.getBean("user1", User.class);
        User user2classTwo = context.getBean("user2", User.class);

        Assertions.assertSame(user1classOne, user1classTwo);
        Assertions.assertSame(user2classOne, user2classTwo);

        // Pada log akan terlihat bawhwa bean pada RenameBeanConfiguration class tidak dibuat.
        // Jadi objek user1classOne == user1classTwo, padahal seharusnya kedua objek itu berbeda,
        // karena didapat dari configuration class yang berbeda.

        // Ini disebabkan karena nama bean atau nama method yang digunakan untuk membuat bean itu sama. Jadinya tabrakan.
        // Solusi supaya tidak terjadi tabrakan seperti itu adalah
        // dengan memberikan nama bean pada Annotation @Bean di method pembuat bean nya.
        // @Bean("namaBean")

        // Bean tidak akan sama dikarenakan nama bean sudah di set,
        // jadi tidak menggunakan default name yaitu dari nama method nya
        User helmiFromClass1 = context.getBean("helmi", User.class);
        User muzakirFromClass2 = context.getBean("muzakir", User.class);
        Assertions.assertNotSame(helmiFromClass1, muzakirFromClass2);

    }
}

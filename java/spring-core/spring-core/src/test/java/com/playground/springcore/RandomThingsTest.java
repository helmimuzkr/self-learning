package com.playground.springcore;

import com.playground.springcore.common.domain.User;
import com.playground.springcore.randomthings.ClassUtil;
import org.junit.jupiter.api.Test;

public class RandomThingsTest {


    @Test
    void testRandomThings() {
        User user = new User();
        System.out.println(ClassUtil.getClassName(user));

        try {
            User newUser = ClassUtil.getInstance(User.class);
            System.out.println("Instance of User -> " + newUser);
        } catch (Exception e) {
            System.out.println("Error : " + e);
        }
    }

}

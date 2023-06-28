package com.myplayground.annotation;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * @Retention untuk memberitahu anotasi akan disimpan dimana.
 * Retention policy RUNTIME -> anotasi akan disimpan di class hasil compile
 * dan pada saat aplikasi dijalankan informasinya dapat diambil.
 *
 * @Target untuk memberitahu anotasi akan digunakan dimana.
 * Element type TYPE -> anotasi bisa digunakan pada class
 * Element type FIELD -> anotasi bisa digunakan pada field
 * Element type METHOD -> anotasi bisa digunakan pada method
 * dan sebagainya.
 */

@Retention(RetentionPolicy.RUNTIME)
@Target({ElementType.TYPE})
public @interface AnnotationCustom {

    String name();

    // Ini adalah contoh field dengan array type
    String[] fields();

    // Kalau dikasih default, berarti field dari annotation ini adalah option
    // bisa diisi kosong ketika digunakan.
    String description() default "ini adalah default description dan annotation retention adalah RUNTIME";
}

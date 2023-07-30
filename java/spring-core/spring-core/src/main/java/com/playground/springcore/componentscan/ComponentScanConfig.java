package com.playground.springcore.componentscan;

import org.springframework.context.annotation.ComponentScan;

// Akan melakukan scan ke path package yang diberikan.

// Jika di dalam package tersebut ada class dengan anotasi @Configration maka akan dibuatkan dengan method pembuatnya,
// Jika ada class @Component maka class tersebut akan langsung diregistrasikan sebagai bean.
// Contoh:
// @Configuration ada di package com.playground.springcore.common.config
// @Component ada di package com.playground.springcore.common.component

// Jika di dalam package tersebut ada package lagi misal config.userrepositoryconfig
// maka package userrepositoryconfig juga akan ikut ter scan.
@ComponentScan(basePackages = {
        "com.playground.springcore.common.config",
        "com.playground.springcore.common.component"
})
public class ComponentScanConfig {
}

package com.playground.springcore.componentscan;

import org.springframework.context.annotation.ComponentScan;

// Package yang di scan tidak ada, maka bean tidak akan ditemukan ketika diambil dari application context/IOC
@ComponentScan(basePackages = {
        "com.playground.springcore.common.packagenotfound"
})
public class ComponentScanConfigError {
}

package com.example.cg.controller;

import com.example.cg.service.UserService;

/**
 * 用户表 控制器
 *
 * @Author melon
 * @Date 2024-03-30 11:30:39
 */
@RestController
@RequestMapping("/user")
public class UserController {
    private UserService userService;
}

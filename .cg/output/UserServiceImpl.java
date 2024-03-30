package com.example.cg.service.impl;

import com.example.cg.dao.domain.User;
import com.example.cg.dao.mapper.UserMapper;
import com.example.cg.service.UserService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;

/**
 * 用户表 业务实现类
 *
 * @Author melon
 * @Date 2024-03-30 11:30:39
 */
@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {
}

package com.example.cg.dao.domain;

import java.util.Date;

import com.baomidou.mybatisplus.annotation.tableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

/**
* user
*
* @Author melon
* @Date 2024-03-27 11:38:48
*/
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@TableName("user")
@ApiModel(value="user表实体类", description="user")
public class User {
    @ApiModelProperty(value = "主键")
    @TableField("id")
    private Long id;

    @ApiModelProperty(value = "姓名")
    @TableField("name")
    private String name;

    @ApiModelProperty(value = "年龄")
    @TableField("age")
    private Integer age;

    @ApiModelProperty(value = "生日")
    @TableField("birth")
    private Date birth;

    @ApiModelProperty(value = "创建时间")
    @TableField("gmt_create")
    private Date gmtCreate;

    @ApiModelProperty(value = "修改时间")
    @TableField("gmt_modify")
    private Date gmtModify;

    @ApiModelProperty(value = "删除标记")
    @TableField("deleted")
    private Byte deleted;
}
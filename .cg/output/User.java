package com.sample.domain;

import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class User {
    /** 主键 */
    private Integer id;
    /** 姓名 */
    private String name;
    /** 年龄 */
    private Integer age;
    /** 生日 */
    private Date birth;
    /** 创建时间 */
    private Date gmtCreate;
    /** 修改时间 */
    private Date gmtModify;
    /** 删除标记 */
    private Boolean deleted;
}
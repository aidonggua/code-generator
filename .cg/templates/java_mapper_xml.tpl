<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="{{config "base-package"}}.{{config "module"}}.{{refs "JavaMapper" "sub-package"}}.{{table "name" | camelCase | title}}{{refs "JavaMapper" "class-postfix"}}">
    <resultMap id="BaseResultMap" type="{{config "base-package"}}.{{config "module"}}.{{refs "JavaEntity" "sub-package"}}.{{table "name" | camelCase | title}}">
    {{range columns -}}
        {{"    "}}<id column="{{.name}}" jdbcType="{{dbToJDBC .type}}" property="{{camelCase .name}}" />
    {{end -}}
    </resultMap>

    <sql id="Base_Column_List">
    {{"    " -}}
    {{range $i,$v := columns -}}
        {{if ne $i 0}},{{end}}{{.name -}}
    {{end}}
    </sql>
</mapper>
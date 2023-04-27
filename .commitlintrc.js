/**
 * type 有关commit的枚举
 * 
 * 1、feat：新功能（feature）
 * 2、fix：修补bug
 * 3、docs：文档（documentation）
 * 4、style： 格式（不影响代码运行的变动）
 * 5、refactor：重构（即不是新增功能，也不是修改bug的代码变动）
 * 6、test：增加测试
 * 7、chore：构建过程或辅助工具的变动
 * 8、revert：回滚到上一个版本
*/

module.exports = {
    extends: ["@commitlint/config-conventional"],
    rules: {
        'type-enum': [
            2,
            'always',
            [
                'feat', 'fix', 'docs', 'style', 'refactor', 'test', 'chore', 'revert',
            ]
        ],
        'type-case': [0],
        'type-empty': [0],
        'scope-empty': [0],
        'scope-case': [0],
        'subject-full-stop': [0, 'never'],
        'subject-case': [0, 'never'],
        'header-max-length': [0, 'always', 72]
    }
};
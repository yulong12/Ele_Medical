class PersonTest {
    public static class Person {
        private String name;
        private int age;

        public void setName(String name) {
            if (name.length() > 6 || name.length() < 2) {
                System.out.println("您设置的人名不符合要求");
                return;
            } else {
                this.name = name;
            }
        }

        public String getName() {
            return this.name;
        }

        public void setAge(int age) {
            if (age > 100 || age < 0) {
                System.out.println("您设置的年龄不合法");
                return;
            } else {
                this.age = age;
            }
        }

        public int getAge() {
            return this.age;
        }
    }

    public static void main(String[] args) {
        Person p = new Person();
        // p.age=1000;
        p.setAge(1000);
        System.out.println("未能设置age成员变量时：" + p.getAge());
        p.setAge(30);
        System.out.println("成功设置age成员变量后：" + p.getAge());
        p.setName("李刚");
        System.out.println("成功设置name成员变量后：" + p.getName());
    }
}
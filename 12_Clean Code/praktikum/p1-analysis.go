package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct {
	t []user
}

func (u userservice) getallusers() []user {
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}

// analysis result:
// 1. the name of structs and function is lack of readability.
//    this can be improved by using the common naming conventions,
//    like:
//     - the structs name and (will be used as public) function
//       must be named with pascal case;
// 2. the userservice seems a redundant struct, since it was a same
//    content like []user type. alternatively, it can be replaced by
//    using direct typing aliasing instead of creating a new struct.
//    it can be written like: `type userservice []user`;
// 3. sub-variable name of userservice struct didn't describe this
//    struct correctfully. the next-maintainer may can't directly
//    recognize what data contained here when the `t`-named variable
//    declared. hence, the code that using this struct can't use this
//    code cleanly. it can be improved by using more desciptive name
//    like `userGroup`, or it can be removed if point 2 is applied;
// 4. struct-managed object isn't using a descriptive name, maybe it
//    will confuse the next-code manager. it can be improved with naming
//    the variable like `users`, and point 1 must be applied first in
//    order to avoid duplication;
// 5. the variable `r` in `getuserbyid` method, doesn't describe what
//    data will be contained here. alternatively, it can be renamed as
//    `user` or `userData`. But, the point 1 must be applied first in
//    order to avoid duplication;
// 6. struct `user` contains unnecessary data to define, like `username`
//    and `password` it can be removed since they don't used in the
//    entire project.

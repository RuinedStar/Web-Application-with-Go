<html>
	<head>
	<title></title>
	</head>
	<body>
		<form action="/login" method="post">
			Username:<input type="text" name="username"></br>
			Password:<input type="password" name="password"></br>
			Chinese Name:<input type="text" name="chinese"></br>
			Email:<input type="text" name="email"></br>
			Favorite fruit:<select name="fruit">
				<option value="apple">apple</option>
				<option value="pear">pear</option>
				<option value="banana">banana</option>
			</select></br>
			Interest:	
				<input type="checkbox" name="interest" value="football">Football
				<input type="checkbox" name="interest" value="basketball">Basketball
				<input type="checkbox" name="interest" value="tennis">Tennis</br>
			Age:<input type="text" name="age"></br>
			Gender:	
				<input type="radio" name="gender" value="1">Male
				<input type="radio" name="gender" value="2">Female
				<input type="radio" name="gender" value="3">Other</br>
			<input type="submit" value="Login">
		</form>
	</body>
</html>
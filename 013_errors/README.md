## Error handling

* `error` is a builtin type. It is an interface with just one method `Error() string`
* Errors are values and they can be anything.
* Keep the happy path inline and only add errors in the if statement.

## Naming Conventions

* Error variables must start with Err<name>. 
* Custom error types must end with <name>Error.

## Comparing Errors

* Compare errors to nil.
* Don't compare error values with one another.
  - See how errors.New() returns the pointer
  - Custom errors must always be pointer receiver

## General Rule

* IF your custom error types have any one of these four methods 
  - Temporary() 
  - Timeout() 
  - NotFound()
  - NotAuthorized() 
 Then have the custom error type have un-exported fields.

 * The method returning the custom error type or error must always
   return the generic error interface and not the custom error type
   See custom_error_find_the_bug.go and custom_error_fix_the_bug.go

   *NOTE* Note that the 'nil' always takes the value of the type.
   i.e. nil of that particular type. So

   ```
   func fail() ([]byte, *customError) {
       return nil, nil
   }
   ```

   The second nil above takes on the `customError` type.

   ```
   +--------------+
   | *customError |
   +--------------+
   | nil          |
   +--------------+
   ```
   The above is not `nil` of `var err error` which would be different.

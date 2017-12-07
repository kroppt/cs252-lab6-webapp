import {
  FormGroup,
  FormControl,
  Validators,
  ValidatorFn,
  AbstractControl
} from '@angular/forms';

/** Passwords do not match */
export function passwordMatchValidator(pass: FormControl): ValidatorFn {
  return (control: AbstractControl): {[key: string]: any} => {
    const nonmatch = pass.value !== control.value;
    return nonmatch ? {'nomatch': {value: control.value}} : null;
  };
}

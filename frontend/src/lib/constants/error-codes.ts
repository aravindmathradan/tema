export enum ServerErrorCodes {
	ENOTFOUND = "not_found",
	EINTERNAL = "internal",
	EMETHODNOTALLOWED = "method_not_allowed",
	EFAILEDVALIDATION = "validation_failed",
	EEDITCONFLICT = "edit_conflict",
	ERATELIMITEXCEEDED = "rate_limit_exceeded",
	EINVALIDCREDENTIALS = "invalid_credentials",
	EINVALIDTOKEN = "invalid_token",
	EBADREQUEST = "bad_request",
	EAUTHREQUIRED = "authentication_required",
	EINACTIVEACCOUNT = "inactive_account",
	ENOTPERMITTED = "not_permitted",
}

export enum ServerErrorSubCodes {
	EBLANKFIELD = "blank_field",
	EVALUENOTPERMITTED = "value_not_permitted",
	EMAXCHARS = "max_chars",
	EMINCHARS = "min_chars",
	EINVALIDEMAIL = "invalid_email",
	EEMAILALREADYEXISTS = "email_already_exists",
	EINVALIDTOKEN = "invalid_token",
	EINVALIDFILTER = "invalid_page_filter",
	ENOTFOUND = "not_found",
	EACCOUNTINACTIVE = "inactive_account",
	EALREADYACTIVE = "already_active",
}

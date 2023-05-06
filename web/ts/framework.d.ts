declare function saveAsFile(textContent: string, filename: string): void;
declare function checkDropdownValue(windowType: "edit" | "create", dropdownType: "gender" | "religion" | "civilstatus"): string;
declare function getDropdownElementIndex(dropdownType: "gender" | "religion" | "civilstatus", dropdownValue: string): string;

export { saveAsFile, checkDropdownValue, getDropdownElementIndex };
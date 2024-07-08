import { ThemedText } from "@/components/ThemedText";
import { Link } from "expo-router";
import { View } from "react-native";

export default function RootIndex() {
  return (
    <View>
      <ThemedText>RootIndex</ThemedText>
      <Link href="/1">
        <ThemedText type="link">Go to Miles!</ThemedText>
      </Link>
      <Link href="/1/food">
        <ThemedText type="link">Go to Food!</ThemedText>
      </Link>
      <Link href="/1/sleep">
        <ThemedText type="link">Go to Sleep!</ThemedText>
      </Link>
      <Link href="/1/add-event">
        <ThemedText type="link">Add event!</ThemedText>
      </Link>
    </View>
  );
}
